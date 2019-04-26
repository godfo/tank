package rest

import "tank/code/config"

/**
 * 系统的所有访问记录均记录在此
 */
type Footprint struct {
	Base
	UserUuid string `json:"userUuid" gorm:"type:char(36)"`
	Ip       string `json:"ip" gorm:"type:varchar(128) not null;index:idx_dt"`
	Host     string `json:"host" gorm:"type:varchar(45) not null"`
	Uri      string `json:"uri" gorm:"type:varchar(255) not null"`
	Params   string `json:"params" gorm:"type:text"`
	Cost     int64  `json:"cost" gorm:"type:bigint(20) not null;default:0"`
	Success  bool   `json:"success" gorm:"type:tinyint(1) not null;default:0"`
}

// set File's table name to be `profiles`
func (this *Footprint) TableName() string {
	return config.TABLE_PREFIX + "footprint"
}