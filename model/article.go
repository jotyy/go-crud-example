package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model

	Title      string `json:"title"`
	CoverUrl   string `json:"cover_url"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

func (article *Article) GetArticles(limit int) ([]*Article, error) {
	var articles []*Article
	err := DB.Model(&Article{}).Limit(limit).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (article *Article) AddArticle() (*Article, error) {
	err := DB.Model(&Article{}).Create(&article).Error
	if err != nil {
		return &Article{}, err
	}
	return article, nil
}
