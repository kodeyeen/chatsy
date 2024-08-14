<template>
    <div class="flex gap-x-[32px]">
        <div class="shrink-0 flex items-center w-[120px]">
            <slot name="label"></slot>
        </div>

        <div class="grow">
            <slot name="entry"></slot>
        </div>
    </div>
</template>
