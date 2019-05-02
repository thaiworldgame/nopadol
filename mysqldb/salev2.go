package mysqldb

import "github.com/mrtomyum/nopadol/sales"

func (repo *salesRepository) FindBankNpRepo() ([]sales.BankModel, error) {

	sql := `select id,company_id,branch_id,code,
	name,active_status from bank`

	Bank := []sales.BankModel{}
	sqls, err := repo.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for sqls.Next() {
		Banks := sales.BankModel{}
		err = sqls.Scan(&Banks.Id, &Banks.CompanyId, &Banks.BranchId, &Banks.Code,
			&Banks.Name, &Banks.ActiveStatus)
		if err != nil {
			return nil, err
		}
		Bank = append(Bank, Banks)
	}

	return Bank, nil
}

func (repo *salesRepository) FindBankBookRepo() ([]sales.BankBookModel, error) {

	sql := `SELECT id,company_id,branch_id,book_no,
	name,account_type,bank_code,branch_code,
	account_code,balance_amount,active_status 
	from bank_book`

	Bank := []sales.BankBookModel{}
	sqls, err := repo.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for sqls.Next() {
		Banks := sales.BankBookModel{}
		err = sqls.Scan(&Banks.Id, &Banks.CompanyId, &Banks.BranchId, &Banks.BookNo, &Banks.Name,
			&Banks.AccountType, &Banks.BankCode, &Banks.BranchCode, &Banks.AccountCode, &Banks.BalanceAmount,
			&Banks.ActiveStatus)
		if err != nil {
			return nil, err
		}
		Bank = append(Bank, Banks)
	}

	return Bank, nil
}

func (repo *salesRepository) FindBankBranchRepo() ([]sales.BankBranchModel, error) {
	sql := `select id,company_id,
	branch_id,bank_branch_code,
	bank_branch_name,active_status
	from bank_branch`

	Bank := []sales.BankBranchModel{}
	sqls, err := repo.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for sqls.Next() {
		Banks := sales.BankBranchModel{}
		err = sqls.Scan(&Banks.Id, &Banks.CompanyId,
			&Banks.BranchId, &Banks.BankBranchCode, &Banks.BankBranchName,
			&Banks.ActiveStatus)
		if err != nil {
			return nil, err
		}
		Bank = append(Bank, Banks)
	}

	return Bank, nil
}

func (repo *salesRepository) FineDepartmentRepo() ([]sales.FineDepartmentModel, error) {

	sql := `select id, company_id, branch_id, code, name, active_status
	from Department`

	Cp := []sales.FineDepartmentModel{}
	sqls, err := repo.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for sqls.Next() {
		Cps := sales.FineDepartmentModel{}
		err = sqls.Scan(&Cps.Id, &Cps.CompanyId,
			&Cps.BranchId, &Cps.Code, &Cps.Name,
			&Cps.ActiveStatus)
		if err != nil {
			return nil, err
		}
		Cp = append(Cp, Cps)
	}
	return Cp, nil
}
