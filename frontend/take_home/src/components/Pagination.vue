<template>
    <!-- The pagination -->
    <div class="ui container">
        <div class="ui pagination menu">
            <a class="active item"> 1 </a>
            <div class="disabled item"> ...  </div>
            <a class="item"> 10 </a>
            <a class="item"> 11 </a>
            <a class="item"> 12 </a>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'Pagination',
        data() {
            return {
                // The current page
                currentPage: 1,
                // The number of items per page
                perPage: 10,
                // The total number of items
                totalItems: 0,
                // The total number of pages
                totalPages: 0,
                // The number of pages to show
                visiblePages: 5,
                // The number of pages to show before and after the current page
                visiblePagesHalf: 2,
                // The number of pages to show before the current page
                visiblePagesBefore: 0,
                // The number of pages to show after the current page
                visiblePagesAfter: 0,
                // The number of pages to show before the current page
                visiblePagesBeforeCurrent: 0,
            }
        }
    }
</script>