!PHONY: start
start:
	make -j 2 runB runF

!PHONY: runB
runB:
	cd chat-app-backend && go run main.go

!PHONY: runF
runF:
	cd chat-app-frontend && pnpm dev

!PHONY: infra
infra:
	cd chat-app-backend && docker compose up -d

!PHONY: teardown
teardown:
	cd chat-app-backend && docker compose down -v