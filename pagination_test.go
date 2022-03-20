package abstraction

import (
	"reflect"
	"testing"
)

func TestPagination_NewPaginationInfo(t *testing.T) {
	type fields struct {
		Page     *int
		PageSize *int
		OrderBy  *string
		Order    *string
	}
	type args struct {
		data interface{}
	}

	p := NewPagination()
	data := make([]*fields, 11)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  *PaginationInfo
	}{
		{
			name: "should success",
			fields: fields{
				Page:     p.Page,
				PageSize: p.PageSize,
				OrderBy:  p.OrderBy,
				Order:    p.Order,
			},
			args: args{
				data: data,
			},
			want: data[:*p.PageSize],
			want1: &PaginationInfo{
				Pagination: p,
				More:       true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pagination := Pagination{
				Page:     tt.fields.Page,
				PageSize: tt.fields.PageSize,
				OrderBy:  tt.fields.OrderBy,
				Order:    tt.fields.Order,
			}
			got, got1 := pagination.NewPaginationInfo(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaginationInfo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewPaginationInfo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
