package test1

var theTableNamePrefix string

////////////////////////////////////////////////////////////////////////////////

// TableName ...
func (Table1) TableName() string {
	return theTableNamePrefix + "_t1"
}

// TableName ...
func (Table2) TableName() string {
	return theTableNamePrefix + "_t2"
}

// TableName ...
func (Table3) TableName() string {
	return theTableNamePrefix + "_t3"
}
