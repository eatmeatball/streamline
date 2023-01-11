// Generated from java-escape by ANTLR 4.11.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link FuncParser}.
 */
public interface FuncListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link FuncParser#primitiveType}.
	 * @param ctx the parse tree
	 */
	void enterPrimitiveType(FuncParser.PrimitiveTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link FuncParser#primitiveType}.
	 * @param ctx the parse tree
	 */
	void exitPrimitiveType(FuncParser.PrimitiveTypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Number}
	 * labeled alternative in {@link FuncParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterNumber(FuncParser.NumberContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Number}
	 * labeled alternative in {@link FuncParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitNumber(FuncParser.NumberContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link FuncParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterMulDiv(FuncParser.MulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link FuncParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitMulDiv(FuncParser.MulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link FuncParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterAddSub(FuncParser.AddSubContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link FuncParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitAddSub(FuncParser.AddSubContext ctx);
	/**
	 * Enter a parse tree produced by {@link FuncParser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterFunctionDeclaration(FuncParser.FunctionDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link FuncParser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitFunctionDeclaration(FuncParser.FunctionDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link FuncParser#classDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterClassDeclaration(FuncParser.ClassDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link FuncParser#classDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitClassDeclaration(FuncParser.ClassDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link FuncParser#start}.
	 * @param ctx the parse tree
	 */
	void enterStart(FuncParser.StartContext ctx);
	/**
	 * Exit a parse tree produced by {@link FuncParser#start}.
	 * @param ctx the parse tree
	 */
	void exitStart(FuncParser.StartContext ctx);
}