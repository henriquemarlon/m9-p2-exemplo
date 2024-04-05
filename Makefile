# DEFAULT VARIABLES
START_LOG = @echo "======================================================= START OF LOG ========================================================="
END_LOG = @echo "======================================================== END OF LOG =========================================================="

tests-kafka:
	@echo "Running the tests for Kafka"
	@docker compose \
		-f ./build/compose.yaml \
		up simulation-mqtt --build -d
	@go test ./test/integration_messaging_test.go -coverprofile=./tools/coverage_sheet.md -v
	@docker compose \
		-f ./build/compose.yaml \
		down simulation-mqtt

tests-mqtt:
	@echo "Running the tests for MQTT"
	@docker compose \
		-f ./build/compose.yaml \
		up simulation-mqtt --build -d
	@go test ./test/integration_mqtt_test.go -coverprofile=./tools/coverage_sheet.md -v
	@docker compose \
		-f ./build/compose.yaml \
		down simulation

.PHONY: env
env:
	$(START_LOG)
	cp ./config/.env.tmpl ./config/.env
	$(END_LOG)
	
.PHONY: mockup
mockup:
	$(START_LOG)
	@docker compose \
		-f ./build/compose.yaml \
		up mockup --build
	$(END_LOG)

.PHONY: run
run:
	$(START_LOG)
	@docker compose \
		-f ./build/compose.yaml \
		--env-file ./config/.env \
		up simulation consumer --build
	$(END_LOG)

