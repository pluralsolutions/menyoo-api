package store

import (
	"github.com/jinzhu/gorm"
	"github.com/plural-solutions/menyoo-api/schema"
)

type EvaluationStore struct {
	*gorm.DB
}

func NewEvaluationStore(db *gorm.DB) *EvaluationStore {
	return &EvaluationStore{db}
}

func (d *EvaluationStore) CreateEvaluation(evaluation *schema.Evaluation) error {
	return d.Where(evaluation).FirstOrCreate(&evaluation).Error
}
