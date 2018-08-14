# command 结构


 ```go
    // Command 是一个命令，适用于您的应用程序。例如。 'go run ...' - 'run'是命令。 Cobra要求您将用法和描述定义为命令定义的一部分，以确保可用性。
    type Command struct {
        // 使用是一行使用消息。.
        Use string
    
        // 别名是一组别名，可用于代替使用中的第一个单词.
        Aliases []string
    
        // SuggestFor是一个命令名数组，建议使用此命令 - 类似于别名但仅建议。
        SuggestFor []string
    
        // Short是“帮助”输出中显示的简短描述.
        Short string
    
        // Long是'help <this-command>'输出中显示的长消息.
        Long string
    
        // 是如何使用该命令的示例.
        Example string
    
        // ValidArgs是bash完成中接受的所有有效非标志参数的列表
        ValidArgs []string
    
        // Expected arguments
        Args PositionalArgs
    
        // ArgAliases是ValidArgs的别名列表。在bash完成时不向用户建议，但如果手动输入则接受.
        ArgAliases []string
    
        // BashCompletionFunction 是bash自动完成生成器使用的自定义函数。
        BashCompletionFunction string
    
        // 如果不推荐使用此命令，则不推荐使用定义，并且在使用时应打印此字符串.
        Deprecated string
    
        // 隐藏定义，如果此命令被隐藏，并且不应显示在可用命令列表中.
        Hidden bool
    
        // 注释是应用程序可用于标识或分组命令的键/值对.
        Annotations map[string]string
    
        // 版本定义此命令的版本。如果此值为非空且命令未定义“version”标志，则将在命令中添加“version”布尔标志，如果指定，将打印“Version”变量的内容.
        Version string
    
        // The *Run functions are executed in the following order:
        //   * PersistentPreRun()
        //   * PreRun()
        //   * Run()
        //   * PostRun()
        //   * PersistentPostRun()
        // All functions get the same args, the arguments after the command name.
        //
        // PersistentPreRun: children of this command will inherit and execute.
        PersistentPreRun func(cmd *Command, args []string)
        // PersistentPreRunE: PersistentPreRun but returns an error.
        PersistentPreRunE func(cmd *Command, args []string) error
        // PreRun: children of this command will not inherit.
        PreRun func(cmd *Command, args []string)
        // PreRunE: PreRun but returns an error.
        PreRunE func(cmd *Command, args []string) error
        // Run: 通常是实际的工作功能。大多数命令只会实现这一点.
        Run func(cmd *Command, args []string)
        // RunE: Run but returns an error.
        RunE func(cmd *Command, args []string) error
        // PostRun: run after the Run command.
        PostRun func(cmd *Command, args []string)
        // PostRunE: PostRun but returns an error.
        PostRunE func(cmd *Command, args []string) error
        // PersistentPostRun: children of this command will inherit and execute after PostRun.
        PersistentPostRun func(cmd *Command, args []string)
        // PersistentPostRunE: PersistentPostRun but returns an error.
        PersistentPostRunE func(cmd *Command, args []string) error
    
        // SilenceErrors是下游安静错误的一种选择。
        SilenceErrors bool
    
        // SilenceUsage是在发生错误时静默使用的选项。
        SilenceUsage bool
    
        // DisableFlagParsing禁用标志解析​​。如果这是真的，所有标志将作为参数传递给命令。
        DisableFlagParsing bool
    
        // DisableAutoGenTag定义，如果生成gen标签（“由spf13 / cobra生成的自动...”）将通过生成此命令的文档来打印。
        DisableAutoGenTag bool
    
        // 在打印帮助或生成文档时，DisableFlagsInUseLine将禁用在命令的使用行中添加[flags]
        DisableFlagsInUseLine bool
    
        // DisableSuggestions disables the suggestions based on Levenshtein distance
        // that go along with 'unknown command' messages.
        DisableSuggestions bool
        // SuggestionsMinimumDistance defines minimum levenshtein distance to display suggestions.
        // Must be > 0.
        SuggestionsMinimumDistance int
    
        // TraverseChildren parses flags on all parents before executing child command.
        TraverseChildren bool
    
        // commands is the list of commands supported by this program.
        commands []*Command
        // parent is a parent command for this command.
        parent *Command
        // Max lengths of commands' string lengths for use in padding.
        commandsMaxUseLen         int
        commandsMaxCommandPathLen int
        commandsMaxNameLen        int
        // commandsAreSorted defines, if command slice are sorted or not.
        commandsAreSorted bool
        // commandCalledAs is the name or alias value used to call this command.
        commandCalledAs struct {
            name   string
            called bool
        }
    
        // args is actual args parsed from flags.
        args []string
        // flagErrorBuf contains all error messages from pflag.
        flagErrorBuf *bytes.Buffer
        // flags is full set of flags.
        flags *flag.FlagSet
        // pflags contains persistent flags.
        pflags *flag.FlagSet
        // lflags contains local flags.
        lflags *flag.FlagSet
        // iflags contains inherited flags.
        iflags *flag.FlagSet
        // parentsPflags is all persistent flags of cmd's parents.
        parentsPflags *flag.FlagSet
        // globNormFunc is the global normalization function
        // that we can use on every pflag set and children commands
        globNormFunc func(f *flag.FlagSet, name string) flag.NormalizedName
    
        // output is an output writer defined by user.
        output io.Writer
        // usageFunc is usage func defined by user.
        usageFunc func(*Command) error
        // usageTemplate is usage template defined by user.
        usageTemplate string
        // flagErrorFunc is func defined by user and it's called when the parsing of
        // flags returns an error.
        flagErrorFunc func(*Command, error) error
        // helpTemplate is help template defined by user.
        helpTemplate string
        // helpFunc is help func defined by user.
        helpFunc func(*Command, []string)
        // helpCommand is command with usage 'help'. If it's not defined by user,
        // cobra uses default help command.
        helpCommand *Command
        // versionTemplate is the version template defined by user.
        versionTemplate string
    }
```