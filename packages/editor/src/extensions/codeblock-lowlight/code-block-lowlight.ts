import * as lowlight from 'lowlight/lib/core'
import CodeBlock, { CodeBlockOptions } from '@tiptap/extension-code-block'
import { LowlightPlugin } from './lowlight-plugin'

export interface CodeBlockLowlightOptions extends CodeBlockOptions {
  lowlight: any,
  defaultLanguage: string | null | undefined,
}

export const CodeBlockLowlight = CodeBlock.extend<CodeBlockLowlightOptions>({
  addOptions() {
    return {
      ...this.parent?.(),
      lowlight: (lowlight as any).lowlight,
      defaultLanguage: null,
    }
  },

  addProseMirrorPlugins() {
    return [
      ...this.parent?.() || [],
      LowlightPlugin({
        name: this.name,
        lowlight: this.options.lowlight,
        defaultLanguage: this.options.defaultLanguage,
      }),
    ]
  },
})