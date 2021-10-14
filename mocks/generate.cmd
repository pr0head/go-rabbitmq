# Use a Vectra Mockery for autogenerate mocks over interfaces
# https://github.com/vektra/mockery

mockery --name=ConsumerInterface --dir=../ --output=.
mockery --name=PublisherInterface --dir=../ --output=.
