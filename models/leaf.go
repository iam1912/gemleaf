package models

import (
	"github.com/iam1912/gemseries/gemleaf/logger"
	"github.com/iam1912/gemseries/gemleaf/utils"
	"gorm.io/gorm"
)

type Leaf struct {
	BizTag      string `gorm:"primaryKey"`
	MaxID       uint64 `gorm:"size:20;default:1"`
	Step        int    `gorm:"size:11"`
	Description string
	UpdatedAt   int64
}

func LeafCreate(db *gorm.DB, tag string, desc string, maxID uint64, step int) error {
	leaf := &Leaf{
		BizTag:      tag,
		MaxID:       maxID,
		Step:        step,
		Description: desc,
	}
	if err := db.Create(&leaf).Error; err != nil {
		logger.Errorf("%s is failed create error:%s", tag, err.Error())
		return err
	}
	logger.Infof("%s is success create\n", tag)
	return nil
}

func LeafGet(db *gorm.DB, tag string) (*Leaf, error) {
	leaf := &Leaf{}
	if err := db.Where("biz_tag = ?", tag).First(&leaf).Error; err != nil {
		return nil, err
	}
	return leaf, nil
}

func LeafUpdate(db *gorm.DB, tag string) error {
	if err := db.Exec("UPDATE leafs SET max_id = max_id + step WHERE biz_tag = ?", tag).Error; err != nil {
		return err
	}
	return nil
}

func LeafNextSegment(db *gorm.DB, tag string) (*Leaf, error) {
	var (
		leaf *Leaf
		err  error
	)
	err = utils.Transaction(db, func(tx *gorm.DB) error {
		if err := LeafUpdate(tx, tag); err != nil {
			return err
		}
		leaf, err = LeafGet(tx, tag)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return leaf, nil
}
