require 'flowdock'

module Logging
  module Appenders
    # Factory for the Flowdock appender.
    def self.flowdock(*args)
      fail ArgumentError, '::Logging::Appenders::Flowdock needs a name as first argument.' if args.empty?
      ::Logging::Appenders::Flowdock.new(*args)
    end

    # This class provides an Appender that can write to a Flowdock service over UDP.
    class Flowdock < ::Logging::Appender
      include Buffering

      # We use a background flusher instead, as the connection to Flowdock can be quite slow:
      DEFAULT_OPTS = {
        flush_period: 1,
      }
      attr_reader :api_token, :flow_user, :flow

      # Creates a new Flowdock Appender that will use the given host and port
      # as the Flowdock server destination.
      #
      # @param name [String] Stream ID to differentiate in the Flowdock server
      # @param api_token [String]
      # @param flow_user [String]
      def initialize(name, opts = {})
        @api_token = opts[:api_token]
        @flow_user = opts[:flow_user]

        fail ArgumentError, 'Empty api_token and flow_user is not appropriate' unless api_token && !api_token.empty? && flow_user && !flow_user.empty?

        # Initialise flow var:
        flow

        super
        configure_buffering(DEFAULT_OPTS.merge(opts))
      end

      def flow
        @flow ||= connect(api_token, flow_user)
      end

      def close(*args)
        super
        @flow = nil

        self
      end

      private

      def connect(host, port)
        ::Flowdock::Flow.new(api_token: api_token, external_user_name: flow_user)
      end

      def canonical_write(str)
        return self if @flow.nil?

        str = str.force_encoding(encoding) if encoding && str.encoding != encoding
        timestamp = Time.now.to_s
        message = "#{name} - #{timestamp}:\n#{str}"
        flow.push_to_chat(content: message)

        self

      rescue StandardError => err
        self.level = :off
        ::Logging.log_internal { "appender #{name.inspect} has been disabled" }
        ::Logging.log_internal_error(err)
      end
    end
  end
end
