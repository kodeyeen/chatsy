<script lang="ts" setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useVModel } from '@vueuse/core'

const props = defineProps({
    modelValue: {
        type: String,
    },

    rows: {
        type: Number,
        default: 1,
    },

    maxlength: {
        type: Number,
    },

    placeholder: {
        type: String,
        default: '',
    },

    name: {
        type: String,
    },

    required: {
        type: Boolean,
        default: false,
    },
})

const emit = defineEmits(['update:modelValue'])
const model = useVModel(props, 'modelValue', emit, {
    defaultValue: '',
})

const isLimited = computed(() => !!props.maxlength)
const charNumber = ref(0)
const counter = ref<any>(null)
const field = ref(null)

const autoresize = (fieldEl: any) => {
    setTimeout(() => {
        fieldEl.style.height = 'auto'
        fieldEl.style.height = `${fieldEl.scrollHeight}px`
    }, 20)
}

const updateCharCounter = (fieldEl: any) => {
    charNumber.value = fieldEl.value.length

    counter.value?.classList.toggle('error', charNumber.value === props.maxlength)
}

const onInput = (event: any) => {
    const fieldEl = event.target
    autoresize(fieldEl)

    if (isLimited.value) {
        updateCharCounter(fieldEl)
    }
}

onMounted(() => {
    autoresize(field.value)
})

watch(
    () => model.value,
    () => {
        autoresize(field.value)
    },
)
</script>

<template>
    <div class="textarea" :class="{ limited: isLimited }">
        <textarea
            ref="field"
            class="textarea__field"
            :rows="rows"
            :maxlength="maxlength"
            :placeholder="placeholder"
            :name="name"
            v-model="model"
            :required="required"
            @input="onInput"
        ></textarea>

        <div v-if="isLimited" ref="counter" class="textarea__counter">
            <span>{{ charNumber }}</span
            >/<span>{{ maxlength }}</span>
        </div>
    </div>
</template>

<style>
.textarea {
    @apply relative
    flex
    flex-col;
}

.textarea__field {
    @apply grow
    block
    w-full
    px-[11px]
    py-[9px]
    border
    rounded-lg
    outline-none
    text-primary-brand-primary
    resize-none
    transition-colors
  
    disabled:hover:border-primary-seattle-30
    disabled:bg-primary-brand-wildSand
    disabled:text-primary-seattle-60
  
    placeholder:focus:opacity-0
  
    cursor-auto
    overflow-y-hidden;
}

.textarea.limited .textarea__field {
    @apply pb-[30px];
}

.textarea.xl .textarea__field,
.textarea.l .textarea__field {
    @apply text-l-long-16;
}

.textarea.m .textarea__field,
.textarea.s .textarea__field {
    @apply text-m-14;
}

.textarea.xs .textarea__field {
    @apply text-s-12;
}

.textarea.filled .textarea__field {
    @apply bg-primary-brand-wildSand
    border-primary-brand-wildSand
    placeholder:text-primary-seattle-100
    hover:border-primary-london-120
    focus:border-primary-seattle-140
    hover:bg-primary-london-100
    focus:bg-primary-brand-white;
}

.textarea.outlined .textarea__field {
    @apply bg-primary-brand-white
    border-primary-seattle-30
    placeholder:text-primary-seattle-60
    hover:border-primary-seattle-140
    focus:border-primary-seattle-140;
}

.textarea.filled.error .textarea__field {
    @apply bg-primary-accentOpacity-7;
}

.textarea.outlined.error .textarea__field {
    @apply border-utilitarian-moscow-100;
}

.textarea__counter {
    @apply absolute
    right-[11px]
    bottom-[9px]
    pointer-events-none
    text-[14px]
    text-primary-seattle-60
    leading-[20px];
}

.textarea.xl .textarea__counter,
.textarea.l .textarea__counter {
    font-size: 14px;
    line-height: 20px;
}

.textarea.m .textarea__counter,
.textarea.s .textarea__counter {
    font-size: 12px;
    line-height: 16px;
}

.textarea.xs .textarea__counter {
    font-size: 12px;
    line-height: 16px;
}

.textarea__counter.error {
    @apply text-utilitarian-moscow-100;
}

.textarea__field::-webkit-scrollbar {
    width: 14px;
}

/* textarea::-webkit-scrollbar-track {
  
  } */

.textarea__field::-webkit-scrollbar-thumb {
    border: 4px solid rgba(0, 0, 0, 0);
    background-clip: padding-box;
    border-radius: 9999px;
    background-color: #aaaaaa;
}
</style>
