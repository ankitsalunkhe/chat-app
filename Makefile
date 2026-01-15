!PHONY: start
start:
	make -j 2 runB runF

!PHONY: runB
runB:
	cd chat-app-backend && go run main.go

!PHONY: runF
runF:
	cd chat-app-frontend && pnpm dev