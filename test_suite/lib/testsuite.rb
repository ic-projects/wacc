
require 'fileutils'
require 'open3'

require 'json'

require 'utils.rb'

class TestSuite

  public

  def initialize(test_cases_path, wacc_compile_script)
    @test_cases_path       = File.absolute_path(test_cases_path)
    @wacc_compile_script   = File.absolute_path(wacc_compile_script)

    @compile    = File.absolute_path(File.join(wacc_compile_script,"compile"))
    @refCompile = File.absolute_path(File.join(wacc_compile_script,"refCompile"))

    sources

  end

  # recomputes the sources and stores them
  def sources
    @sources = Dir.glob(File.join(@test_cases_path,"**","*.wacc")).sort
    @sources
  end

  # takes a block argument which is called back by each result
  def run_tests(&block)

    @sources.each do |source|
      result = test_source(File.absolute_path(source))
      block.call({ :source => source, :result => result })
    end

  end

  def run_test(source, &block)
    if @sources.member?(source) then
      result = test_source(source)
      block.call({:source => source, :result => result})
    end
  end

  private

  def test_source(source)
    compile_results = run_frontend(source)

    { :source => source,
      :compile => compile_results,
      :passed => compile_results[:passed]
    }
  end

  def run_frontend(source)

    unless File.exists?(@compile) then
      return { :result => :error,
               :passed => false,
               :message => "compile script cannot be found"
             }
    end

    unless File.exists?(@refCompile) then
      return { :result => :error,
               :passed => false,
               :message => "refCompile script cannot be found"
             }
    end

    ast_file = File.dirname(source) + "/" +  File.basename(source, ".wacc") + ".ast"
    expected = File.read(ast_file)
    run_results = Utils.run3("/usr/bin/timeout",
                       ["--kill-after=5", "3", @compile, "-t", source],
                       nil, 1024 * 1024 * 100)

    passed = (expected == run_results[:stdout])

    return { :result => :ran,
             :passed => passed,
             :run_results => run_results,
             :expected    => { :stdout => expected }
           }
  end

end
