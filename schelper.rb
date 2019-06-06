class Schelper < Formula
  desc "Analyze log files from Sauce Connect"
  homepage "https://github.com/mdsauce/schelper"
  url "https://github.com/mdsauce/schelper/archive/v1.0.2.tar.gz"
  sha256 "bcb427f47ca22a7eed382ac1919eb4d3642653c7b35f9ed1b38a70842f69ef9d"
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
