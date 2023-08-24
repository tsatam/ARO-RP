# Copyright (c) Microsoft Corporation.
# Licensed under the Apache License 2.0.

import urllib3

from azext_aro.vendored_sdks.azure.mgmt.redhatopenshift.v2023_04_01 import AzureRedHatOpenShiftClient
from azure.cli.core.commands.client_factory import get_mgmt_service_client


def cf_aro(cli_ctx, *_):

    opt_args = {}

    client = get_mgmt_service_client(cli_ctx,
                                     AzureRedHatOpenShiftClient,
                                     **opt_args)

    return client
