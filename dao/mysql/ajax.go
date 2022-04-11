package mysql

// AjaxChangeStatus 使用 Ajax 更新指定表的指定id的status
func AjaxChangeStatus(id, table, field string) (err error) {
	// update TABLE_NAME set FIELD_NAME = ABS(FIELD_NAME - 1) where id = ? AND is_deleted=0;

	if RowsAffected := db.Exec("update "+table+" set "+field+" = ABS("+field+" - 1) where id = ? AND is_deleted=0;", id).RowsAffected; RowsAffected != 1 {
		return ErrAjaxChangeStatus
	}
	return nil
}
