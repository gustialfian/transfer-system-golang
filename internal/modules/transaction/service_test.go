package transaction

import (
	"context"
	"fmt"
	"testing"

	"github.com/gustialfian/transfer-system-golang/internal/modules/account"
)

func TestTransactionService_Create(t *testing.T) {
	type fields struct {
		repo        TransactionRepo
		accountRepo account.AccountRepo
	}
	type args struct {
		ctx  context.Context
		data TransactionCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "error - invalid amount",
			fields: fields{
				repo:        &fakeTransactionRepo{},
				accountRepo: &fakeAccountRepo{},
			},
			args: args{
				ctx: t.Context(),
				data: TransactionCreate{
					SourceAccountId:      1,
					DestinationAccountId: 2,
					Amount:               "aaa",
				},
			},
			wantErr: true,
		},
		{
			name: "error - negative amount",
			fields: fields{
				repo:        &fakeTransactionRepo{},
				accountRepo: &fakeAccountRepo{},
			},
			args: args{
				ctx: t.Context(),
				data: TransactionCreate{
					SourceAccountId:      1,
					DestinationAccountId: 2,
					Amount:               "-1",
				},
			},
			wantErr: true,
		},
		{
			name: "error - same source and destination account",
			fields: fields{
				repo:        &fakeTransactionRepo{},
				accountRepo: &fakeAccountRepo{},
			},
			args: args{
				ctx: t.Context(),
				data: TransactionCreate{
					SourceAccountId:      1,
					DestinationAccountId: 1,
					Amount:               "1",
				},
			},
			wantErr: true,
		},
		{
			name: "error - account not found",
			fields: fields{
				repo: &fakeTransactionRepo{},
				accountRepo: &fakeAccountRepo{
					ByIdFunc: func(ctx context.Context, accountId int) (account.AccountRow, error) {
						return account.AccountRow{}, fmt.Errorf("test-error")
					},
				},
			},
			args: args{
				ctx: t.Context(),
				data: TransactionCreate{
					SourceAccountId:      1,
					DestinationAccountId: 2,
					Amount:               "1",
				},
			},
			wantErr: true,
		},
		{
			name: "error - account balance not enough",
			fields: fields{
				repo: &fakeTransactionRepo{},
				accountRepo: &fakeAccountRepo{
					ByIdFunc: func(ctx context.Context, accountId int) (account.AccountRow, error) {
						return account.AccountRow{
							AccountId:    1,
							Balance:      100_000,
							ScaleBalance: 5,
						}, nil
					},
				},
			},
			args: args{
				ctx: t.Context(),
				data: TransactionCreate{
					SourceAccountId:      1,
					DestinationAccountId: 2,
					Amount:               "10",
				},
			},
			wantErr: true,
		},
		{
			name: "error - fail update account balance",
			fields: fields{
				repo: &fakeTransactionRepo{},
				accountRepo: &fakeAccountRepo{
					ByIdFunc: func(ctx context.Context, accountId int) (account.AccountRow, error) {
						return account.AccountRow{
							AccountId:    1,
							Balance:      1_000_000,
							ScaleBalance: 5,
						}, nil
					},
					UpdateBalanceFunc: func(ctx context.Context, params account.AccountUpdateBalanceParams) error {
						return fmt.Errorf("test-error")
					},
				},
			},
			args: args{
				ctx: t.Context(),
				data: TransactionCreate{
					SourceAccountId:      1,
					DestinationAccountId: 2,
					Amount:               "1",
				},
			},
			wantErr: true,
		},
		{
			name: "error - fail create transaction",
			fields: fields{
				repo: &fakeTransactionRepo{
					CreateFunc: func(ctx context.Context, data TransactionCreateParams) error {
						return fmt.Errorf("test-error")
					},
				},
				accountRepo: &fakeAccountRepo{
					ByIdFunc: func(ctx context.Context, accountId int) (account.AccountRow, error) {
						return account.AccountRow{
							AccountId:    1,
							Balance:      1_000_000,
							ScaleBalance: 5,
						}, nil
					},
					UpdateBalanceFunc: func(ctx context.Context, params account.AccountUpdateBalanceParams) error {
						return nil
					},
				},
			},
			args: args{
				ctx: t.Context(),
				data: TransactionCreate{
					SourceAccountId:      1,
					DestinationAccountId: 2,
					Amount:               "1",
				},
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				repo: &fakeTransactionRepo{
					CreateFunc: func(ctx context.Context, data TransactionCreateParams) error {
						return nil
					},
				},
				accountRepo: &fakeAccountRepo{
					ByIdFunc: func(ctx context.Context, accountId int) (account.AccountRow, error) {
						return account.AccountRow{
							AccountId:    1,
							Balance:      1_000_000,
							ScaleBalance: 5,
						}, nil
					},
					UpdateBalanceFunc: func(ctx context.Context, params account.AccountUpdateBalanceParams) error {
						return nil
					},
				},
			},
			args: args{
				ctx: t.Context(),
				data: TransactionCreate{
					SourceAccountId:      1,
					DestinationAccountId: 2,
					Amount:               "1",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTransactionService(tt.fields.repo, tt.fields.accountRepo)
			if err := svc.Create(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("TransactionService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type fakeTransactionRepo struct {
	CreateFunc func(ctx context.Context, data TransactionCreateParams) error
}

func (f *fakeTransactionRepo) Create(ctx context.Context, data TransactionCreateParams) error {
	return f.CreateFunc(ctx, data)
}

type fakeAccountRepo struct {
	CreateFunc        func(ctx context.Context, data account.AccountCreateParams) error
	ByIdFunc          func(ctx context.Context, accountId int) (account.AccountRow, error)
	UpdateBalanceFunc func(ctx context.Context, params account.AccountUpdateBalanceParams) error
}

func (f *fakeAccountRepo) Create(ctx context.Context, data account.AccountCreateParams) error {
	return f.CreateFunc(ctx, data)
}

func (f *fakeAccountRepo) ById(ctx context.Context, accountId int) (account.AccountRow, error) {
	return f.ByIdFunc(ctx, accountId)
}

func (f *fakeAccountRepo) UpdateBalance(ctx context.Context, params account.AccountUpdateBalanceParams) error {
	return f.UpdateBalanceFunc(ctx, params)
}
