package account

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestAccountService_Create(t *testing.T) {
	type fields struct {
		repo AccountRepo

		isTigerBeetleOn bool
		tigerbeetleRepo AccountTBRepo
	}
	type args struct {
		ctx  context.Context
		data AccountCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "error - invalid intial balance alpha",
			fields: fields{repo: &fakeAccountRepo{
				CreateFunc: func(ctx context.Context, data AccountCreateParams) error { return nil },
			}},
			args: args{
				ctx:  t.Context(),
				data: AccountCreate{AccountId: 1, InitialBalance: "aaa"},
			},
			wantErr: true,
		},
		{
			name: "error - invalid intial balance negative",
			fields: fields{repo: &fakeAccountRepo{
				CreateFunc: func(ctx context.Context, data AccountCreateParams) error { return nil },
			}},
			args: args{
				ctx:  t.Context(),
				data: AccountCreate{AccountId: 1, InitialBalance: "-1"},
			},
			wantErr: true,
		},
		{
			name: "error - db fail",
			fields: fields{repo: &fakeAccountRepo{
				CreateFunc: func(ctx context.Context, data AccountCreateParams) error { return fmt.Errorf("test-error") },
			}},
			args: args{
				ctx:  t.Context(),
				data: AccountCreate{AccountId: 1, InitialBalance: "100.23344"},
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{repo: &fakeAccountRepo{
				CreateFunc: func(ctx context.Context, data AccountCreateParams) error { return nil },
			}},
			args: args{
				ctx:  t.Context(),
				data: AccountCreate{AccountId: 1, InitialBalance: "100.23344"},
			},
			wantErr: false,
		},
		{
			name: "success - tigerbeetle",
			fields: fields{
				repo: &fakeAccountRepo{
					CreateFunc: func(ctx context.Context, data AccountCreateParams) error { return nil },
				},
				isTigerBeetleOn: true,
				tigerbeetleRepo: &fakeAccountTBRepo{
					CreateAccountFunc:     func(accountId int) error { return nil },
					CreateTransactionFunc: func(debitAccountId, creditAccountId, amount int) error { return nil },
				},
			},
			args: args{
				ctx:  t.Context(),
				data: AccountCreate{AccountId: 1, InitialBalance: "100.23344"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewAccountService(tt.fields.repo, tt.fields.tigerbeetleRepo, tt.fields.isTigerBeetleOn)
			if err := svc.Create(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("AccountService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountService_ById(t *testing.T) {
	type fields struct {
		repo AccountRepo

		isTigerBeetleOn bool
		tigerbeetleRepo AccountTBRepo
	}
	type args struct {
		ctx       context.Context
		accountId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		{
			name: "error - db fail",
			fields: fields{repo: &fakeAccountRepo{
				ByIdFunc: func(ctx context.Context, accountId int) (AccountRow, error) {
					return AccountRow{}, fmt.Errorf("test-error")
				},
			}},
			args: args{
				ctx:       t.Context(),
				accountId: 1,
			},
			want:    Account{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{repo: &fakeAccountRepo{
				ByIdFunc: func(ctx context.Context, accountId int) (AccountRow, error) {
					return AccountRow{
						AccountId:    1,
						Balance:      100_000,
						ScaleBalance: 5,
					}, nil
				},
			}},
			args: args{
				ctx:       t.Context(),
				accountId: 1,
			},
			want: Account{
				AccountId:      1,
				InitialBalance: "1.00000",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewAccountService(tt.fields.repo, tt.fields.tigerbeetleRepo, tt.fields.isTigerBeetleOn)
			got, err := svc.ById(tt.args.ctx, tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountService.ById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountService.ById() = %v, want %v", got, tt.want)
			}
		})
	}
}

type fakeAccountRepo struct {
	CreateFunc        func(ctx context.Context, data AccountCreateParams) error
	ByIdFunc          func(ctx context.Context, accountId int) (AccountRow, error)
	UpdateBalanceFunc func(ctx context.Context, params AccountUpdateBalanceParams) error
}

func (f *fakeAccountRepo) Create(ctx context.Context, data AccountCreateParams) error {
	return f.CreateFunc(ctx, data)
}

func (f *fakeAccountRepo) ById(ctx context.Context, accountId int) (AccountRow, error) {
	return f.ByIdFunc(ctx, accountId)
}

func (f *fakeAccountRepo) UpdateBalance(ctx context.Context, params AccountUpdateBalanceParams) error {
	return f.UpdateBalanceFunc(ctx, params)
}

type fakeAccountTBRepo struct {
	CreateAccountFunc     func(accountId int) error
	CreateTransactionFunc func(debitAccountId int, creditAccountId int, amount int) error
}

func (f *fakeAccountTBRepo) CreateAccount(accountId int) error {
	return f.CreateAccountFunc(accountId)
}

func (f *fakeAccountTBRepo) CreateTransaction(debitAccountId int, creditAccountId int, amount int) error {
	return f.CreateTransactionFunc(debitAccountId, creditAccountId, amount)
}
