<script>
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import { pageTitle, appName, hideControls } from "@/stores/app";
    import { addSuccessToast } from "@/stores/toasts";
    import tooltip from "@/actions/tooltip";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import Field from "@/components/base/Field.svelte";
    import SettingsSidebar from "@/components/settings/SettingsSidebar.svelte";
    import ContentSidebar from "./ContentSidebar.svelte";

    let originalFormSettings = {};
    let formSettings = {};
    let isLoading = false;
    let isSaving = false;
    let initialHash = "";

    $: initialHash = JSON.stringify(originalFormSettings);

    $: hasChanges = initialHash != JSON.stringify(formSettings);

    loadSettings();

    async function loadSettings() {
        isLoading = true;

        try {
            const settings = (await ApiClient.settings.getAll()) || {};
            init(settings);
        } catch (err) {
            ApiClient.error(err);
        }

        isLoading = false;
    }

    async function save() {
        if (isSaving || !hasChanges) {
            return;
        }

        isSaving = true;

        try {
            const settings = await ApiClient.settings.update(CommonHelper.filterRedactedProps(formSettings));
            init(settings);
            addSuccessToast("Successfully saved application settings.");
        } catch (err) {
            ApiClient.error(err);
        }

        isSaving = false;
    }

    function init(settings = {}) {
        $appName = settings?.meta?.appName;
        $hideControls = !!settings?.meta?.hideControls;

        formSettings = {
            meta: settings?.meta || {},
        };

        originalFormSettings = JSON.parse(JSON.stringify(formSettings));
    }

    function reset() {
        formSettings = JSON.parse(JSON.stringify(originalFormSettings || {}));
    }
</script>

<ContentSidebar />

<PageWrapper>
    <header class="page-header">
        <nav class="breadcrumbs">
            <div class="breadcrumb-item">Content Management</div>
        </nav>
    </header>

    <div class="wrapper">
        <form class="panel" autocomplete="off" on:submit|preventDefault={save}>
            {#if isLoading}
                <div class="loader" />
            {:else}
                <div class="grid">
                    <div class="col-lg-6">
                        <Field class="form-field required" name="meta.primaryColour" let:uniqueId>
                            <label for={uniqueId}>Primary colour</label>
                            <input
                                type="color"
                                id={uniqueId}
                                required
                                bind:value={formSettings.meta.primaryColour}
                            />
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field required" name="meta.secondaryColour" let:uniqueId>
                            <label for={uniqueId}>Secondary colour</label>
                            <input
                                type="color"
                                id={uniqueId}
                                required
                                bind:value={formSettings.meta.secondaryColour}
                            />
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field required" name="meta.logo" let:uniqueId>
                            <label for={uniqueId}>Secondary colour</label>
                            <input type="file" id={uniqueId} required bind:value={formSettings.meta.logo} />
                        </Field>
                    </div>

                    <Field class="form-field form-field-toggle" name="meta.hideControls" let:uniqueId>
                        <input type="checkbox" id={uniqueId} bind:checked={formSettings.meta.hideControls} />
                        <label for={uniqueId}>
                            <span class="txt">Hide collection create and edit controls</span>
                            <i
                                class="ri-information-line link-hint"
                                use:tooltip={{
                                    text: `This could prevent making accidental schema changes when in production environment.`,
                                    position: "right",
                                }}
                            />
                        </label>
                    </Field>

                    <div class="col-lg-12 flex">
                        <div class="flex-fill" />

                        {#if hasChanges}
                            <button
                                type="button"
                                class="btn btn-transparent btn-hint"
                                disabled={isSaving}
                                on:click={() => reset()}
                            >
                                <span class="txt">Cancel</span>
                            </button>
                        {/if}
                        <button
                            type="submit"
                            class="btn btn-expanded"
                            class:btn-loading={isSaving}
                            disabled={!hasChanges || isSaving}
                            on:click={() => save()}
                        >
                            <span class="txt">Save changes</span>
                        </button>
                    </div>
                </div>
            {/if}
        </form>
    </div>
</PageWrapper>
