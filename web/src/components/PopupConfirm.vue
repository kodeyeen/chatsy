<script lang="ts" setup>
import { usePopupsStore } from '@/stores/popups'
import Button from '@/components/Button.vue'
import Popup from '@/components/Popup.vue'

const props = defineProps<{
    id: string
    title: string
    description?: string
    confirmText: string
    cancelText: string
}>()

const popupsStore = usePopupsStore()

const emit = defineEmits(['confirm', 'cancel'])

const onConfirm = () => {
    popupsStore.hide(props.id)
    emit('confirm')
}

const onCancel = () => {
    popupsStore.hide(props.id)
    emit('cancel')
}
</script>

<template>
    <Popup :id="id" dialogClass="md:max-w-[450px]" contentClass="md:px-[16px] md:py-[20px]">
        <div>
            <div class="text-h5 text-primary-brand-primary" v-text="title"></div>
        </div>

        <div v-if="description" class="mt-[16px]">
            <div v-text="description"></div>
        </div>

        <div class="flex flex-col md:flex-row gap-[8px] mt-[20px]">
            <Button class="grow" type="button" size="l" variant="critical" round @click="onConfirm">
                <span v-text="confirmText"></span>
            </Button>

            <Button class="grow" type="button" size="l" variant="clear" round @click="onCancel">
                <span v-text="cancelText"></span>
            </Button>
        </div>
    </Popup>
</template>
