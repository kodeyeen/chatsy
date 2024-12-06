<script lang="ts" setup>
import { ref, computed, watch } from 'vue'


import { useChatsStore } from '@/stores/chats'
import { usePopupsStore } from '@/stores/popups'


const props = defineProps<{
    filter: string
}>()

const emit = defineEmits(['update:filter'])

const model = computed({
    get() {
        return props.filter
    },

    set(newValue) {
        emit('update:filter', newValue)
    },
})

const popupsStore = usePopupsStore()

const filter = ref('')

watch(filter, (newFilter) => {
    model.value = newFilter
})
</script>

<template>
    <div>
        

        
    </div>
</template>
