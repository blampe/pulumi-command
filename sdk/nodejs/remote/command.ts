// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A command to run on a remote host.
 * The connection is established via ssh.
 */
export class Command extends pulumi.CustomResource {
    /**
     * Get an existing Command resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Command {
        return new Command(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'command:remote:Command';

    /**
     * Returns true if the given object is an instance of Command.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Command {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Command.__pulumiType;
    }

    /**
     * The parameters with which to connect to the remote host
     */
    public readonly connection!: pulumi.Output<outputs.remote.Connection | undefined>;
    /**
     * The command to run on create.
     */
    public readonly create!: pulumi.Output<string | undefined>;
    /**
     * The command to run on delete.
     */
    public readonly delete!: pulumi.Output<string | undefined>;
    /**
     * Additional environment variables available to the command's process.
     */
    public readonly environment!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The standard error of the command's process
     */
    public /*out*/ readonly stderr!: pulumi.Output<string>;
    /**
     * Pass a string to the command's process as standard in
     */
    public readonly stdin!: pulumi.Output<string | undefined>;
    /**
     * The standard output of the command's process
     */
    public /*out*/ readonly stdout!: pulumi.Output<string>;
    /**
     * Trigger replacements on changes to this input.
     */
    public readonly triggers!: pulumi.Output<any[] | undefined>;
    /**
     * The command to run on update, if empty, create will run again.
     */
    public readonly update!: pulumi.Output<string | undefined>;

    /**
     * Create a Command resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CommandArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(inputs.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["create"] = args ? args.create : undefined;
            resourceInputs["delete"] = args ? args.delete : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["stdin"] = args ? args.stdin : undefined;
            resourceInputs["triggers"] = args ? args.triggers : undefined;
            resourceInputs["update"] = args ? args.update : undefined;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
        } else {
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["create"] = undefined /*out*/;
            resourceInputs["delete"] = undefined /*out*/;
            resourceInputs["environment"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdin"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
            resourceInputs["triggers"] = undefined /*out*/;
            resourceInputs["update"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Command.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Command resource.
 */
export interface CommandArgs {
    /**
     * The parameters with which to connect to the remote host.
     */
    connection: pulumi.Input<inputs.remote.ConnectionArgs>;
    /**
     * The command to run on create.
     */
    create?: pulumi.Input<string>;
    /**
     * The command to run on delete.
     */
    delete?: pulumi.Input<string>;
    /**
     * Additional environment variables available to the command's process.
     */
    environment?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Pass a string to the command's process as standard in
     */
    stdin?: pulumi.Input<string>;
    /**
     * Trigger replacements on changes to this input.
     */
    triggers?: pulumi.Input<any[]>;
    /**
     * The command to run on update, if empty, create will run again.
     */
    update?: pulumi.Input<string>;
}
