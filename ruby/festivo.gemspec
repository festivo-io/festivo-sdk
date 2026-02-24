Gem::Specification.new do |spec|
  spec.name          = "festivo"
  spec.version       = "0.2.0"
  spec.authors       = ["Festivo Team"]
  spec.email         = ["support@getfestivo.com"]
  spec.summary       = "Festivo Public Holidays API Ruby SDK"
  spec.description   = "Official Ruby SDK for Festivo Public Holidays API. Access city-level, regional, and global holiday data."
  spec.homepage      = "https://getfestivo.com"
  spec.license       = "MIT"

  spec.files         = Dir["lib/**/*.rb"]
  spec.require_paths = ["lib"]

  spec.add_runtime_dependency "net-http", ">= 0"
  spec.add_runtime_dependency "json", ">= 0"

  spec.metadata["homepage_uri"] = spec.homepage
  spec.metadata["source_code_uri"] = "https://github.com/festivo-io/festivo-sdk"
  spec.metadata["documentation_uri"] = "https://getfestivo.com/docs"
end

