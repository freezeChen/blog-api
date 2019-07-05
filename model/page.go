/*
   @Time : 2019-07-05 14:54
   @Author : frozenChen
   @File : page
   @Software: blog-api
*/
package model

type Page struct {
	Page int
	Size int
}

func (p *Page) GetOffSet() int {
	if p.Page > 1 {
		return (p.Page - 1) * p.Size
	}

	return 0
}
