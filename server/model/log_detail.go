package model

type LogDetail struct {
	ID        uint     `json:"id" db:"id"`
	LogNameID uint     `json:"logNameId" db:"log_name_id"`
	LogName   *LogName `json:"logName" db:"log_names"`
	Version   uint     `json:"version" db:"version"`
}
