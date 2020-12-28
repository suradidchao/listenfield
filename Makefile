## dev-start: Start docker-compose
dev-start:
	@ echo "> Start development environment"
	@ docker-compose up -d --build

seed-data:
	@ echo "> Seeding data"
	@ seed/seed.sh

## dev-stop: Start docker-compose
dev-stop:
	@ echo "> Stop development environment"
	@ docker-compose down
