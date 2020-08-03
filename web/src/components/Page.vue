<template>
  <v-container class="home">
    <v-row
      justify="center"
      no-gutters
    >
      <v-col lg="10">
        <v-card>
          <v-row
            class="px-8 py-2"
          >
            <v-col>
              <h1>{{ title }}</h1>
              <slot />
            </v-col>
          </v-row>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

export default {
  name: 'Page',
  props: {
    title: {
      type: String,
      required: true,
    },
  },
};
</script>
