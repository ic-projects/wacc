
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
    frontend_results = run_frontend(source)
    backend_results = run_backend(source)

    { :source => source,
      :frontend => frontend_results,
      :backend => backend_results,
      :passed => frontend_results[:passed] && backend_results[:passed]
    }
  end

  def run_frontend(source)

    unless File.exists?(@compile) then
      return { :result => :error,
               :passed => false,
               :message => "compile script cannot be found"
             }
    end

    ast_file = File.dirname(source) + "/" +  File.basename(source, ".wacc") + ".ast"
    expected = File.read(ast_file)
    run_results = Utils.run3("/usr/bin/timeout",
                       ["--kill-after=5", "3", @compile, "-t", source],
                       nil, 1024 * 1024 * 100)
    expected_exit = "0"

    if(File.dirname(source) =~ /invalid\/semanticErr/)
      expected_exit = "200"
    elsif(File.dirname(source) =~ /invalid\/syntaxErr/)
      expected_exit = "100"
    end

    if(expected_exit == "100" || expected_exit == "200")
      passed = (run_results[:exit_status].to_s == expected_exit)
    else
      passed = (expected == run_results[:stdout])
    end

    return { :result => :ran,
             :passed => passed,
             :run_results => run_results,
             :expected    => { :stdout => expected, :exit_status => expected_exit}
           }
  end

  def run_backend(source)

    unless File.exists?(@compile) then
      return { :result => :error,
               :passed => false,
               :message => "compile script cannot be found"
             }
    end

    out_file = File.dirname(source) + "/" +  File.basename(source, ".wacc") + ".out"
    expected = File.read(out_file)

    run_results = Utils.run3("/usr/bin/timeout",
                       ["--kill-after=5", "3", @compile, "-x", source],
                       nil, 1024 * 1024 * 100)

    expected_exit = "0"

    if(File.dirname(source) =~ /invalid\/semanticErr/)
      expected_exit = "200"
    elsif(File.dirname(source) =~ /invalid\/syntaxErr/)
      expected_exit = "100"
    end

    if(expected_exit == "100" || expected_exit == "200")
      passed = (run_results[:exit_status].to_s == expected_exit)
    else
      passed = (expected == run_results[:stdout])
    end

    expected_asmFile = File.dirname(source) + "/" +  File.basename(source, ".wacc") + ".s"
    if(File.file?(expected_asmFile))
      expected_asm = File.read(expected_asmFile)
    else
      expected_asm = ""
    end

    asmFile = File.absolute_path(File.join(@wacc_compile_script, File.basename(source, ".wacc") + ".s"))
    puts asmFile
    if(File.file?(asmFile))
      asm = File.read(asmFile)
      removeAssembly = %x(rm #{asmFile})
    else
      asm = ""
    end

    return { :result => :ran,
             :asm => asm,
             :expected_asm => expected_asm,
             :passed => passed,
             :run_results => run_results,
             :expected    => { :stdout => expected, :exit_status => expected_exit}
           }
  end

end
