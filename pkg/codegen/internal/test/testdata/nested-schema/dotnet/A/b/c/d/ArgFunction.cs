// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Example.A/b/c/d
{
    public static class ArgFunction
    {
        public static Task<ArgFunctionResult> InvokeAsync(ArgFunctionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ArgFunctionResult>("example:a/b/c/d:argFunction", args ?? new ArgFunctionArgs(), options.WithDefaults());

        public static Output<ArgFunctionResult> Invoke(ArgFunctionInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ArgFunctionResult>("example:a/b/c/d:argFunction", args ?? new ArgFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class ArgFunctionArgs : Pulumi.InvokeArgs
    {
        [Input("arg1")]
        public string? Arg1 { get; set; }

        public ArgFunctionArgs()
        {
        }
    }

    public sealed class ArgFunctionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arg1")]
        public Input<string>? Arg1 { get; set; }

        public ArgFunctionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ArgFunctionResult
    {
        public readonly string? Result;

        [OutputConstructor]
        private ArgFunctionResult(string? result)
        {
            Result = result;
        }
    }
}