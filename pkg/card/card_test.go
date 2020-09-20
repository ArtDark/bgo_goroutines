package card

import (
	"reflect"
	"testing"
)

func TestSortTransactions(t *testing.T) {

	tests := []struct {
		name string
		args []Transaction
		want []Transaction
	}{
		{
			name: "Сортировка транзакций по сумме от большего к меньшему",
			args: []Transaction{
				{
					Id:     "1",
					Bill:   1_203_91,
					Time:   1592594432,
					MCC:    "5812",
					Status: "Valid",
				},
				{
					Id:     "2",
					Bill:   755_64,
					Time:   1592595432,
					MCC:    "5812",
					Status: "Valid",
				},
				{
					Id:     "3",
					Bill:   125_12,
					Time:   1592596432,
					MCC:    "5931",
					Status: "Valid",
				},
				{
					Id:     "4",
					Bill:   4_563_45,
					Time:   1592614432,
					MCC:    "5931",
					Status: "Valid",
				},
			},
			want: []Transaction{
				{
					Id:     "4",
					Bill:   4_563_45,
					Time:   1592614432,
					MCC:    "5931",
					Status: "Valid",
				},
				{
					Id:     "1",
					Bill:   1_203_91,
					Time:   1592594432,
					MCC:    "5812",
					Status: "Valid",
				},
				{
					Id:     "2",
					Bill:   755_64,
					Time:   1592595432,
					MCC:    "5812",
					Status: "Valid",
				},
				{
					Id:     "3",
					Bill:   125_12,
					Time:   1592596432,
					MCC:    "5931",
					Status: "Valid",
				},
			},
		},
	}
	for _, tt := range tests {
		if SortTransactions(tt.args); !reflect.DeepEqual(tt.args, tt.want) {
			t.Errorf("Sum() = %v, want %v", tt.args, tt.want)
		}
	}
}
