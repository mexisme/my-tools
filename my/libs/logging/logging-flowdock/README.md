# Logging::Flowdock

This is an "Appender" for the [Logging](https://rubygems.org/gems/logging) gem, to allow you to forward Logs directly to an
[Flowdock](https://flowdock.com) Service.

## Installation

Add this line to your application's Gemfile:

```ruby
gem 'logging'
gem 'logging-flowdock'
```

And then execute:

    $ bundle

Or install it yourself as:

    $ gem install logging-flowdock

## Usage


The Logging gem uses [little-plugger](https://rubygems.org/gems/little-plugger) gem to manage
[extending Logger](https://github.com/TwP/logging#extending).
This means it will automatically load and initialise the `logging-flowdock` code as soon as it loads the Logging code.

To configure Logging to use Flowdock, you simply have to add the appender to your configuring of Logging.

e.g. The following will send the "Wootles" log-output to STDERR and the Flowdock service running on the localhost:

```ruby
l = Logging.logger["The Wild Wild Test"]
l.add_appenders(
  Logging.appenders.stderr,
  Logging.appenders.flowdock("Wicky wicky wild", api_token: "AKSDHKSAHDKJAHSDKH-IS-MADE-UP", flow_user: "Bot")
)

l.info "Wootles"
```

You must provide an `api_token` and a `flow_user`.

### Logging Defaults

Sending messages to the Flowdock service can be quite slow, so by default I've set the `flush_period` to 1 second.
This means a `log.info("MESSAGE")` call should return immediately, and a background thread will flush all logging messages to the Flowdock
service once every second.

If you prefer to *not* do this (e.g. thread-unsafe code) you can disable it:

```
[...]

l.add_appenders(
  Logging.appenders.flowdock(name, ..., flush_period: nil)
)
```

## Development

After checking out the repo, run `bin/setup` to install dependencies. Then, run `rake spec` to run the tests. You can also run `bin/console ${API_TOKEN} ${FLOWDOCK_USER}` for an interactive prompt that will allow you to experiment.

To install this gem onto your local machine, run `bundle exec rake install`. To release a new version, update the version number in `version.rb`, and then run `bundle exec rake release`, which will create a git tag for the version, push git commits and tags, and push the `.gem` file to [rubygems.org](https://rubygems.org).

### To Do's

I track major "To Do's" in a [TODO.org](TODO.org) file.

I track context-sensitive and minor ones with `TODO` and `FIXME` comments in the source code.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/mexisme/logging-flowdock. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.


## License

The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
