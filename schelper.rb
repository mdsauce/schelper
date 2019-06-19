class Schelper < Formula
  desc "Analyze log files from Sauce Connect"
  homepage "https://github.com/mdsauce/schelper"
  url "https://github.com/mdsauce/schelper/archive/v1.0.3.tar.gz"
  sha256 "0b59c8211421755f86098a0a621de69ec6385bf0ea8bddbe175bc0dc182594e0"
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
