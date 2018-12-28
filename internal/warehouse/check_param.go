package warehouse

func CheckInWarehouseTableParam(table *InWarehouseTableBody) bool {
	if table.Writer == 0 || table.Material == nil || table.ProjectName == "" || table.TableNum == ""{
		return false
	}
	return true
}

func CheckOutWarehouseTableParam(table *OutWarehouseTableBody) bool {
	if table.Writer == 0 || table.Material == nil || table.TableNum == "" || table.ProjectName == ""{
		return false
	}
	return true
}

func CheckConfirmOutWarehouseParam(table *ConfirmOutWarehouseBody) bool {
	if table.TableID == 0 || table.ConfirmUser == 0 {
		return false
	}
	return true
}
