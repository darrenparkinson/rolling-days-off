<template>
  <div>
    <v-data-table
      :loading="loading"
      loading-text="Loading... Please wait"
      :headers="headers"
      :items="daysOff"
      class="elevation-3"
    >
      <template v-slot:item.day="{ item }">{{ getDay(item.date) }}</template>
      <template v-slot:item.date="{ item }">{{ formatDate(item.date) }}</template>
    </v-data-table>
  </div>
</template>
<script>
export default {
  props: {
    daysOff: Array,
    loading: Boolean
  },
  data() {
    return {
      headers: [
        { text: "Day", align: "left", value: "day" },
        { text: "Date", align: "left", value: "date" },
        { text: "", align: "left", value: "bankHolidayTitle" },
        { text: "", align: "left", value: "bankHolidayNotes" }
      ]
    };
  },
  methods: {
    getDay(date) {
      const options = { weekday: "long" };
      const dayOff = new Date(date);
      return dayOff.toLocaleDateString("en-GB", options);
    },
    formatDate(date) {
      const options = {
        // weekday: "long",
        year: "numeric",
        month: "long",
        day: "2-digit"
      };
      let dayOff = new Date(date);
      return dayOff.toLocaleDateString("en-GB", options);
    }
  }
};
</script>
<style>
tbody tr:nth-of-type(odd) {
  background-color: rgba(0, 0, 0, 0.05);
}
::v-deep .v-data-table__empty-wrapper {
  display: none;
}
</style>