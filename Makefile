gen-user:
	protoc proto/user/user.proto --go_out=plugins=grpc:.
gen-company:
	protoc proto/company/company.proto --go_out=plugins=grpc:.
gen-transaction:
	protoc proto/transaction/transaction.proto --go_out=plugins=grpc:.