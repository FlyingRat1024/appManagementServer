package material

func CheckApplyTableParam(table *ApplyTableBody) bool {
	if table.MaterialID == 0 || table.ApplierID == 0 {
		return false
	}
	return true
}

func CheckReceiveTableParam(table *RecieveTableBody) bool {
	if table.MaterialID == 0 || table.ReciiverID == 0 || table.Num == 0 {
		return false
	}
	return true
}

func CheckBackTableParam(table *BackTableBody) bool {
	if table.TableID == 0 || table.BackerID == 0 || table.Num == 0 || table.MaterialID == 0{
		return false
	}
	return true
}

func CheckCheckTableParam(table *CheckTableBody) bool {
	if table.TableID == 0 || table.CheckerID == 0 || table.Num == 0 || table.MaterialID == 0 {
		return false
	}
	return true
}
