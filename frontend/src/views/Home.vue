<template>
  <v-container fluid>
    <div class="home">
      <RollingDaysOff />
    </div>
  </v-container>
</template>

<script>
// @ is an alias to /src
import RollingDaysOff from "@/components/RollingDaysOff.vue";

export default {
  name: "home",
  components: {
    RollingDaysOff
  }
};
</script>
