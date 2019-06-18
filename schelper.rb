class Schelper < Formula
  desc "Analyze log files from Sauce Connect"
  homepage "https://github.com/mdsauce/schelper"
  url "https://github.com/mdsauce/schelper/archive/v1.0.2.tar.gz"
  sha256 "a225045946c0a2b8efe242c3ccaeb6ecceabb77e71095236e4269fe2299ac592"
  depends_on "go" => :build
  version "1.0.2"

  def install
    system "go", "build", "-o", bin/"schelper", "."
  end

  test do
    system "#{bin}/schelper", --help, ">", "output.txt"
    assert_predicate ./"output.txt", :exist?
  end
end
