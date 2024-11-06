lexer:
	@cd runner && go run main.go lexer

semantic:
	@cd runner && go run main.go semantic

3ac:
	@cd runner && go run main.go 3ac

run:
	@cd runner && go run main.go
