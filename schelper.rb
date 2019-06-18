class Schelper < Formula
  desc "Analyze log files from Sauce Connect"
  homepage "https://github.com/mdsauce/schelper"
  url "https://github.com/mdsauce/schelper/archive/v1.0.3.tar.gz"
  sha256 "8ae9dfb235fa8effa896b1f27e024efb69894419de4eeb913a1dc0578f85a7f9"
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
