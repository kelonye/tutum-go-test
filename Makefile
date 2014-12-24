OPTS=

deploy:
	@$(MAKE) build
	@$(MAKE) push
	@$(MAKE) up

build:
	@$(OPTS) td build

push:
	@$(OPTS) td push

up:
	@$(OPTS) td up

ps:
	@$(OPTS) td ps

.PHONY: ps up push build deploy
