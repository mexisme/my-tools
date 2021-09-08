module Logging
  module Plugins
    # This is plugin-intialisation module used by `little-plugger` to find the code:
    module Flowdock
      extend self

      VERSION = '0.1.0'.freeze unless defined? VERSION

      # Initialiser for `little-plugger`:
      def initialize_flowdock
        require File.expand_path('../../appenders/flowdock', __FILE__)
      end
    end
  end
end
