<script>
    import { link } from "svelte-spa-router";
    import active from "svelte-spa-router/active";
    import PageSidebar from "@/components/base/PageSidebar.svelte";
</script>

<PageSidebar class="settings-sidebar">
    <div class="sidebar-content">
        <div class="sidebar-title">CMS</div>
        <a href="/content" class="sidebar-list-item" use:active={{ path: "/content" }} use:link>
            <i class="ri-home-gear-line" aria-hidden="true" />
            <span class="txt">Content managementuu</span>
        </a>
    </div>
</PageSidebar>
