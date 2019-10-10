package cmd

func (a *AttackManager) getCurrentRunningAttackCount() int {
	return len(a.repository)
}
