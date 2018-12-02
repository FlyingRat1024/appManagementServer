package material

func CheckApplyTableParam(table *ApplyTableBody) bool {
	if table.Material == nil || table.ApplierID == 0 {
		return false
	}
	return true
}

func CheckReceiveTableParam(table *RecieveTableBody) bool {
	if table.Material == nil || table.ReciiverID == 0 {
		return false
	}
	return true
}

func CheckBackTableParam(table *BackTableBody) bool {
	if table.TableID == 0 || table.BackerID == 0 || table.Material == nil {
		return false
	}
	return true
}

func CheckCheckTableParam(table *CheckTableBody) bool {
	if table.TableID == 0 || table.CheckerID == 0 || table.Material == nil {
		return false
	}
	return true
}
