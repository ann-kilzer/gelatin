<template>
  <page title="My cool page" />
</template>

<script>
// @ is an alias to /src
import Page from '@/components/Page.vue';

export default {
  name: 'Home',
  components: {
    Page,
  },
};
</script>
