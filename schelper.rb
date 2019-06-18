class Schelper < Formula
  desc "Analyze log files from Sauce Connect"
  homepage "https://github.com/mdsauce/schelper"
  url "https://github.com/mdsauce/schelper/archive/v1.0.3.tar.gz"
  sha256 "36d01c99ad9dea9abc5c96e3e43e91e21e2d36469ecc65d6f396b2e4d350cb49"
  depends_on "go" => :build
  version "1.0.3"

  def install
    system "go", "build", "-o", bin/"schelper", "."
  end

  test do
    system "#{bin}/schelper", --help, ">", "output.txt"
    assert_predicate ./"output.txt", :exist?
  end
end
