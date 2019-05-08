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

func (repo *salesRepository) FindProductByKey(Keyword string) ([]sales.ProductModal, error) {
	sql := `select a.code as item_code,
	ifnull(c.bar_code,'') as bar_code,
	a.item_name,
	ifnull(a.pic_path1,'') as pic_path_1,
	ifnull(a.stock_type,0) as stock_type,
	ifnull(b.rate1,1) as rate_1,
	ifnull(b.unit_code,'') as unit_code,
	ifnull(d.sale_price_1,0) as sale_price_1,
	ifnull(d.sale_price_2,0) as sale_price_2,
	ifnull(e.wh_code,'') as wh_code,
	ifnull(e.shelf_code,'') as shelf_code,
	ifnull(e.qty,0) as qty,
	ifnull(e.unit_code,'') as stk_unit_code
	from Item a 
	left join ItemRate b on a.code = b.item_code
	left join Barcode c on a.code = c.item_code
	left join Price d on a.code = d.item_code 
	and b.unit_code = d.unit_code
	LEFT JOIN  (select wh_code,item_code,shelf_code,qty,unit_code  from StockLocation GROUP by item_code) e on a.code = e.item_code
	where a.code like concat(?,'%') or item_name like concat(?,'%') GROUP by a.code
	order by a.code LIMIT 100`

	product := []sales.ProductModal{}
	pb, err := repo.db.Query(sql, Keyword, Keyword)
	if err != nil {
		return nil, err
	}
	for pb.Next() {
		products := sales.ProductModal{}
		err = pb.Scan(&products.ItemCode, &products.BarCode, &products.ItemName,
			&products.PicPath1, &products.StockType,
			&products.Rate1, &products.UnitCode,
			&products.SalePrice1, &products.SalePrice2, &products.WHCode,
			&products.ShelfCode, &products.Qty, &products.StkUnitCode)
		if err != nil {
			return nil, err
		}
		// sqlpb := `select ifnull(wh_code,'') as wh_code,ifnull(shelf_code,'') as shelf_code,
		// 	ifnull(b.qty,0) as qty,ifnull(b.unit_code,'') as stk_unit_code
		// 	from Item a left join StockLocation b on a.code = b.item_code
		// 	where a.code = ? order by b.wh_code`
		// subproduct := []sales.SubproductModal{}
		// subpb, err := repo.db.Query(sqlpb, products.ItemCode)
		// if err != nil {
		// 	return nil, err
		// }
		// for subpb.Next() {
		// 	subproducts := sales.SubproductModal{}
		// 	err = subpb.Scan(&subproducts.WHCode, &subproducts.ShelfCode,
		// 		&subproducts.Qty, &subproducts.StkUnitCode)
		// 	subproduct = append(subproduct, subproducts)

		// }
		// for _, subit := range subproduct {
		// 	p := sales.SubproductModal{
		// 		WHCode:      subit.WHCode,
		// 		ShelfCode:   subit.ShelfCode,
		// 		Qty:         subit.Qty,
		// 		StkUnitCode: subit.StkUnitCode,
		// 	}
		// 	products.Sub = append(products.Sub, p)
		// }

		product = append(product, products)
	}
	// for i, pbo := range product {

	// 	subproduct := []sales.SubproductModal{}
	// 	subpb, err := repo.db.Query(sqlpb, pbo.ItemCode)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	for subpb.Next() {
	// 		subproducts := sales.SubproductModal{}
	// 		err = subpb.Scan(&subproducts.WHCode, &subproducts.ShelfCode,
	// 			&subproducts.Qty, &subproducts.StkUnitCode)
	// 		subproduct = append(subproduct, subproducts)

	// 	}
	// 	for _, subit := range subproduct {
	// 		p := sales.SubproductModal{
	// 			WHCode:    subit.WHCode,
	// 			ShelfCode: subit.ShelfCode,

	// 			Qty:         subit.Qty,
	// 			StkUnitCode: subit.StkUnitCode,
	// 		}
	// 		product[i].Sub = append(product[i].Sub, p)
	// 	}

	// }
	return product, nil
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
	from Department `

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
