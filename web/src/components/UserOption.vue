<script lang="ts" setup>
import { computed } from 'vue'

import Button from '@/components/Button.vue'
import DoneCheckIcon from '@/components/icons/DoneCheck.vue'
import PlusIcon from '@/components/icons/Plus.vue'
import UserItem from '@/components/UserItem.vue'

const props = defineProps<{
    user: any,
    modelValue: any,
    disabled?: boolean,
}>()

const emit = defineEmits(['update:modelValue'])

const model = computed({
    get() {
        return props.modelValue
    },

    set(newValue) {
        emit('update:modelValue', newValue)
    },
})

const isSelected = computed(() => {
    return model.value.some((user: any) => user.id === props.user.id)
})

const onSelect = () => {
    if (isSelected.value) {
        model.value = model.value.filter((user: any) => user.id !== props.user.id)
    } else {
        model.value = [...model.value, props.user]
    }
}
</script>

<template>
    <UserItem :user="user">
        <template #end>
            <Button
                class="w-[32px] px-[6px] md:w-auto md:px-[15px] md:gap-x-[6px]"
                type="button"
                size="s"
                :variant="isSelected ? 'accent' : 'clear'"
                :disabled="disabled"
                round
                @click="onSelect"
            >
                <DoneCheckIcon v-if="isSelected" :width="20" :height="20" />
                <PlusIcon v-else :width="20" :height="20" />

                <span class="hidden md:inline">
                    {{ isSelected ? 'Selected' : 'Select' }}
                </span>
            </Button>
        </template>
    </UserItem>
</template>
