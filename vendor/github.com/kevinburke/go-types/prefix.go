package types

import (
	"database/sql/driver"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kevinburke/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// A PrefixUUID stores an additional prefix as part of a UUID type.
type PrefixUUID struct {
	Prefix string
	UUID   uuid.UUID
}

func (u PrefixUUID) String() string {
	return u.Prefix + u.UUID.String()
}

// GenerateUUID generates a UUID with the given prefix.
func GenerateUUID(prefix string) PrefixUUID {
	uid := uuid.NewV4()
	id := PrefixUUID{
		Prefix: prefix,
		UUID:   uid,
	}
	return id
}

// NewPrefixUUID creates a PrefixUUID from the prefix and string uuid. Returns
// an error if uuidstr cannot be parsed as a valid UUID.
func NewPrefixUUID(caboodle string) (PrefixUUID, error) {
	if len(caboodle) < 36 {
		return PrefixUUID{}, fmt.Errorf("types: Could not parse \"%s\" as a UUID with a prefix", caboodle)
	}
	uuidPart := caboodle[len(caboodle)-36:]
	u, err := uuid.FromString(uuidPart)
	if err != nil {
		return PrefixUUID{}, err
	}

	return PrefixUUID{
		Prefix: caboodle[:len(caboodle)-36],
		UUID:   u,
	}, nil
}

func (pu *PrefixUUID) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	p, err := NewPrefixUUID(s)
	if err != nil {
		return err
	}
	*pu = p
	return nil
}

func (pu PrefixUUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(pu.String())
}

// Scan implements the Scanner interface. Note only the UUID gets scanned/set
// here, we can't determine the prefix from the database. `value` should be
// a [16]byte
func (pu *PrefixUUID) Scan(value interface{}) error {
	if value == nil {
		return errors.New("types: cannot scan null into a PrefixUUID")
	}
	var err error
	switch t := value.(type) {
	case []byte:
		if len(t) >= 32 {
			*pu, err = NewPrefixUUID(string(t))
		} else {
			var u uuid.UUID
			u, err = uuid.FromBytes(t)
			pu.UUID = u
		}
	case string:
		*pu, err = NewPrefixUUID(t)
	default:
		return fmt.Errorf("types: can't scan value of unknown type %v into a PrefixUUID", value)
	}
	return err
}

// Value implements the driver.Valuer interface.
func (pu PrefixUUID) Value() (driver.Value, error) {
	// In theory we should be able to send 16 raw bytes to the database
	// and have it encoded as a UUID. However, this requires enabling
	// binary_parameters=yes on the connection string. Instead of that, just
	// pass a string to the database, which is easy to handle.
	return pu.UUID.String(), nil
}

// GetBSON implements the mgo.Getter interface.
func (pu PrefixUUID) GetBSON() (interface{}, error) {
	return bson.Binary{
		Kind: 0x03,
		Data: pu.UUID[:],
	}, nil
}

// SetBSON implements the mgo.Setter interface.
func (pu *PrefixUUID) SetBSON(raw bson.Raw) error {
	// first 4 bytes are int32 LE length
	if len(raw.Data) < 4 {
		return fmt.Errorf("invalid BSON data: too short")
	}
	l := binary.LittleEndian.Uint32(raw.Data[:4])
	var err error
	if l >= 32 {
		// null terminated, so subtract 1
		d := string(raw.Data[4 : len(raw.Data)-1])
		*pu, err = NewPrefixUUID(d)
	} else {
		var u uuid.UUID
		u, err = uuid.FromBytes(raw.Data[4:])
		pu.UUID = u
	}
	return err
}
