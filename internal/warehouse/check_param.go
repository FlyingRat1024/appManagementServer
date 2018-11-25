package warehouse

func CheckInWarehouseTableParam(table *InWarehouseTableBody) bool {
	if table.Name == "" || table.Unit == "" || table.Number == 0 || table.Writer == 0 {
		return false
	}
	return true
}

func CheckOutWarehouseTableParam(table *OutWarehouseTableBody) bool {
	if table.Writer == 0 || table.MaterialID == 0 {
		return false
	}
	return true
}
