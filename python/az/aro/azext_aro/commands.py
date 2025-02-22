# Copyright (c) Microsoft Corporation.
# Licensed under the Apache License 2.0.

from azure.cli.core.commands import CliCommandType
from azext_aro._client_factory import cf_aro
from azext_aro._format import aro_show_table_format
from azext_aro._format import aro_list_table_format
from azext_aro._format import aro_version_table_format
from azext_aro._help import helps  # pylint: disable=unused-import


def load_command_table(self, _):
    aro_sdk = CliCommandType(
        operations_tmpl='azext_aro.vendored_sdks.azure.mgmt.redhatopenshift.v2023_09_04.operations#OpenShiftClustersOperations.{}',  # pylint: disable=line-too-long
        client_factory=cf_aro)

    with self.command_group('aro', aro_sdk, client_factory=cf_aro) as g:
        g.custom_command('create', 'aro_create', supports_no_wait=True)
        g.custom_command('delete', 'aro_delete', supports_no_wait=True, confirmation=True)
        g.custom_command('list', 'aro_list', table_transformer=aro_list_table_format)
        g.custom_show_command('show', 'aro_show', table_transformer=aro_show_table_format)
        g.custom_command('update', 'aro_update', supports_no_wait=True)
        g.wait_command('wait')

        g.custom_command('list-credentials', 'aro_list_credentials')
        g.custom_command('get-admin-kubeconfig', 'aro_get_admin_kubeconfig')

        g.custom_command('get-versions', 'aro_get_versions', table_transformer=aro_version_table_format)

        g.custom_command('validate', 'aro_validate')
