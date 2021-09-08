# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'logging/plugins/flowdock'

Gem::Specification.new do |spec|
  spec.name          = "logging-flowdock"
  spec.version       = Logging::Plugins::Flowdock::VERSION
  spec.authors       = ["mexisme"]
  spec.email         = ["wildjim+dev@kiwinet.org"]

  spec.summary       = 'An appender for the Logging gem that sends messages to a Flowdock channel'
  # spec.description   = %q{TODO: Write a longer description or delete this line.}
  spec.homepage      = 'https://github.com/fairfaxmedia/logging-flowdock'
  spec.license       = "MIT"

  spec.files         = `git ls-files -z`.split("\x0").reject { |f| f.match(%r{^(test|spec|features)/}) }
  spec.bindir        = "exe"
  spec.executables   = spec.files.grep(%r{^exe/}) { |f| File.basename(f) }
  spec.require_paths = ["lib"]

  spec.add_development_dependency "bundler", "~> 1.11"
  spec.add_development_dependency "rake", "~> 10.0"
  spec.add_development_dependency "rspec", "~> 3.3"
  spec.add_development_dependency 'pry-byebug', '~> 3.2'

  spec.add_runtime_dependency 'little-plugger', '~> 1.1'
  spec.add_runtime_dependency 'logging', '~> 2.0'
  spec.add_runtime_dependency 'flowdock'
end
