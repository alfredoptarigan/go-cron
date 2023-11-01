inject:
	@echo "Run wire to inject dependencies"
	@cd pkg/injection && wire
	@echo "Dependencies injected"