class Schelper < Formula
  desc "Analyze log files from Sauce Connect"
  homepage "https://github.com/mdsauce/schelper"
  url "https://github.com/mdsauce/schelper/archive/v1.0.3.tar.gz"
  sha256 "0cc6f91e40699578f3183b07dc5768d0e3482996a40233e4baac3c776a2c1dbd"
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
